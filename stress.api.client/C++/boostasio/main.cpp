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

// g++ -std=c++17 -Wall -Wextra -I /usr/include main.cpp -o server -L /usr/lib -lboost_system -lboost_thread -lpthread -lboost_coroutine -lboost_context -lboost_chrono -lboost_date_time
// sudo apt-get install libboost-all-dev
void handle_request(http::request<http::string_body>& req, http::response<http::string_body>& res) {
    res.version(req.version());
    res.keep_alive(false);

    if (req.method() == http::verb::get) {
        res.result(http::status::ok);
        res.set(http::field::content_type, "application/json");
        res.body() = R"({"name": "John Doe", "cpf": "12345678901", "telefone": "555-1234"})";
        res.content_length(res.body().size());
    } else {
        res.result(http::status::bad_request);
        res.set(http::field::content_type, "text/plain");
        res.body() = "Invalid request.";
        res.content_length(res.body().size());
    }
}

void session(tcp::socket socket) {
    beast::flat_buffer buffer;

    while (true) {
        http::request<http::string_body> req;
        http::read(socket, buffer, req);

        http::response<http::string_body> res;
        handle_request(req, res);

        http::write(socket, res);

        if (res.need_eof()) {
            break;
        }
    }

    socket.shutdown(tcp::socket::shutdown_send);
}

int main() {
    try {
        auto const address = asio::ip::make_address("0.0.0.0");
        auto const port = static_cast<unsigned short>(8080);

        asio::io_context ioc{1};

        tcp::acceptor acceptor{ioc, {address, port}};
        while (true) {
            tcp::socket socket{ioc};
            acceptor.accept(socket);

            std::thread{[sock = std::move(socket)]() mutable { session(std::move(sock)); }}.detach();
        }
    } catch (std::exception const& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 1;
    }

    return 0;
}
