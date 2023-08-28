#include <iostream>
#include <boost/beast/core.hpp>
#include <boost/beast/http.hpp>
#include <boost/asio/ip/tcp.hpp>
#include <boost/asio.hpp>
#include <boost/config.hpp>
#include <cstdlib>
#include <ctime>
#include <memory>
#include <string>

using tcp = boost::asio::ip::tcp;
namespace http = boost::beast::http;

class session : public std::enable_shared_from_this<session> {
    tcp::socket socket_;
    boost::beast::flat_buffer buffer_;
    http::request<http::string_body> req_;
    http::response<http::string_body> res_;

public:
    explicit session(tcp::socket socket)
        : socket_(std::move(socket)) {}

    void run() {
        read_request();
    }

private:
    void read_request() {
        auto self = shared_from_this();
        http::async_read(socket_, buffer_, req_,
            [self](boost::beast::error_code ec, std::size_t) {
                if (!ec) self->handle_request();
            });
    }

    void handle_request() {
        if (req_.method() == http::verb::get && req_.target() == "/v1/client/get") {
            res_.version(req_.version());
            res_.keep_alive(req_.keep_alive());
            res_.set(http::field::server, "Beast");
            res_.set(http::field::content_type, "application/json");

            res_.body() = R"({"name": "Jeffotoni", "cpf": "484.487.789-52", "telefone": "(11) 98765-87654"})";
            res_.content_length(res_.body().size());
            res_.result(http::status::ok);
        } else {
            res_.result(http::status::not_found);
            res_.body() = "Not Found";
            res_.content_length(res_.body().size());
            res_.set(http::field::content_type, "application/json");
        }
        write_response();
    }

    void write_response() {
        auto self = shared_from_this();
        http::async_write(socket_, res_,
            [self](boost::beast::error_code ec, std::size_t) {
                if (!ec) self->socket_.shutdown(tcp::socket::shutdown_send, ec);
            });
    }
};

class server {
    tcp::acceptor acceptor_;
    tcp::socket socket_;

public:
    server(boost::asio::io_context& ioc, const tcp::endpoint& endpoint)
        : acceptor_(ioc, endpoint), socket_(ioc) {
        accept();
    }

private:
    void accept() {
        acceptor_.async_accept(socket_, [&](boost::beast::error_code ec) {
            if (!ec) std::make_shared<session>(std::move(socket_))->run();
            accept();
        });
    }
};

int main() {
    try {
        auto const address = boost::asio::ip::make_address("0.0.0.0");
        auto const port = static_cast<unsigned short>(8080);

        std::cout << "Run Server C++ Port: 0.0.0.0:8080" << std::endl;

        boost::asio::io_context ioc{1};
        server srv{ioc, tcp::endpoint{address, port}};
        ioc.run();

    } catch (std::exception const& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 

        1;
    }

    return 0;
}
