#include <boost/asio.hpp>
#include <boost/beast.hpp>
#include <boost/beast/http.hpp>
#include <boost/beast/version.hpp>
#include <iostream>
#include <memory>
#include <string>
#include <thread>
#include <vector>

#include "AsyncHttpClient.h"

namespace http = boost::beast::http;
using tcp = boost::asio::ip::tcp;

// g++ -std=c++17 -Wall -Wextra -I /usr/include main.cpp AsyncHttpClient.cpp -o server -L /usr/lib -lboost_system -lboost_thread -lpthread -lboost_coroutine -lboost_context -lboost_chrono -lboost_date_time

class session : public std::enable_shared_from_this<session> {
public:
    explicit session(tcp::socket socket) : socket_(std::move(socket)) {}

    void start() {
        read_request();
    }

private:
    void read_request() {
        auto self = shared_from_this();
        http::async_read(socket_, buffer_, req_,
                         [self](boost::beast::error_code ec, std::size_t bytes_transferred) {
                             boost::ignore_unused(bytes_transferred);
                             if (!ec) {
                                 self->handle_request();
                             }
                         });
    }

    void handle_request() {
        if (req_.method() == http::verb::get && req_.target() == "/v1/client/get") {
            auto self = shared_from_this();
            boost::asio::io_context ioc;
            std::shared_ptr<AsyncHttpClient> client = std::make_shared<AsyncHttpClient>(
                ioc, "localhost", "3000", "/v1/customer/get",
                [self](const std::string& response_body) {
                    self->res_.result(http::status::ok);
                    self->res_.set(http::field::content_type, "application/json");
                    self->res_.body() = response_body;
                    self->res_.prepare_payload();
                    self->write_response();
                });

            client->start();
            ioc.run();
        } else {
            res_.result(http::status::not_found);
            res_.set(http::field::content_type, "text/plain");
            res_.body() = "The resource '" + std::string(req_.target()) + "' was not found.";
            res_.prepare_payload();
            write_response();
        }
    }

    void write_response() {
        auto self = shared_from_this();
        http::async_write(socket_, res_,
                          [self](boost::beast::error_code ec, std::size_t bytes_transferred) {
                              boost::ignore_unused(bytes_transferred);
                              if (!ec) {
                                  boost::system::error_code ignored_ec;
                                  self->socket_.shutdown(tcp::socket::shutdown_send, ignored_ec);
                              }
                          });
    }

    tcp::socket socket_;
    boost::beast::flat_buffer buffer_;
    http::request<http::string_body> req_;
    http::response<http::string_body> res_;
};

class server {
public:
    server(boost::asio::io_context& ioc, const tcp::endpoint& endpoint) : acceptor_(ioc, endpoint) {}

    void run() {
        do_accept();
    }

private:
    void do_accept() {
        acceptor_.async_accept([this](boost::beast::error_code ec, tcp::socket socket) {
            if (!ec) {
                std::make_shared<session>(std::move(socket))->start();
            }
            do_accept();
        });
    }

    tcp::acceptor acceptor_;
};

int main() {
    try {
        auto const address = boost::asio::ip::make_address("0.0.0.0");
        auto const port = static_cast<unsigned short>(8080);
        auto const threads = 10;
        
        std::cout << "Run Server C++ Port: 0.0.0.0:8080" << std::endl;

        boost::asio::io_context ioc{threads};
        server srv{ioc, tcp::endpoint{address, port}};
        srv.run();

        std::vector<std::thread> v;
        v.reserve(threads - 1);
        for (auto i = threads - 1; i > 0; --i) {
            v.emplace_back([&ioc] { ioc.run(); });
        }
        ioc.run();

        for (auto& t : v) {
            t.join();
        }

    } catch (std::exception const& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 1;
    }

    return 0;
}
