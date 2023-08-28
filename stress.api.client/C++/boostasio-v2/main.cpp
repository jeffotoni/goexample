#include <boost/asio.hpp>
#include <boost/beast.hpp>
#include <boost/beast/http.hpp>
#include <iostream>
#include <string>
#include <thread>

namespace asio = boost::asio;
namespace beast = boost::beast;
namespace http = beast::http;

using tcp = asio::ip::tcp;

std::string fetch_customer_data(asio::io_context& ioc) {
    tcp::resolver resolver{ioc};
    beast::tcp_stream stream{ioc};

    auto const results = resolver.resolve("localhost", "3000");

    stream.connect(results);

    http::request<http::empty_body> req{http::verb::get, "/v1/customer/get", 11};
    req.set(http::field::host, "localhost");
    req.set(http::field::user_agent, "cpp-client");

    http::write(stream, req);

    beast::flat_buffer buffer;
    http::response<http::string_body> res;

    http::read(stream, buffer, res);

    stream.socket().shutdown(tcp::socket::shutdown_both);

    return res.body();
}

void handle_request(http::request<http::string_body>& req, http::response<http::string_body>& res, asio::io_context& ioc) {
    res.version(req.version());
    res.keep_alive(false);

   // if (req.method() == http::verb::get) {
    if (req.method() == http::verb::get && req.target() == "/v1/client/get") {
        res.result(http::status::ok);
        res.set(http::field::content_type, "application/json");
        res.body() = fetch_customer_data(ioc);
        res.content_length(res.body().size());
    } else {
        res.result(http::status::bad_request);
        res.set(http::field::content_type, "application/json");
        res.body() = "{\"msg\":\"Invalid request.\"}";
        res.content_length(res.body().size());
    }
}

void session(tcp::socket socket, asio::io_context& ioc) {
    beast::flat_buffer buffer;

    while (true) {
        http::request<http::string_body> req;
        http::read(socket, buffer, req);

        http::response<http::string_body> res;
        handle_request(req, res, ioc);

        http::write(socket, res);

        if (res.need_eof()) {
            break;
        }
    }

    socket.shutdown(tcp::socket::shutdown_send);
}

// https://github.com/boostorg/boost
// g++ -std=c++17 -Wall -Wextra -I /usr/include main.cpp -o server -L /usr/lib -lboost_system -lboost_thread -lpthread -lboost_coroutine -lboost_context -lboost_chrono -lboost_date_time
int main() {
    try {
        auto const address = asio::ip::make_address("0.0.0.0");
        auto const port = static_cast<unsigned short>(8080);
        
        std::cout << "Run Server C++ Port: 0.0.0.0:8080" << std::endl;

        asio::io_context ioc{1};

        tcp::acceptor acceptor{ioc, {address, port}};
        while (true) {
            tcp::socket socket{ioc};
            acceptor.accept(socket);

            std::thread{[sock = std::move(socket), &ioc]() mutable { session(std::move(sock), ioc); }}.detach();
        }
    } catch (std::exception const& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 1;
    }
    return 0;
}
