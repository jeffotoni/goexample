#include "AsyncHttpClient.h"

AsyncHttpClient::AsyncHttpClient(boost::asio::io_context& ioc, const std::string& host,
                                 const std::string& port, const std::string& target,
                                 std::function<void(const std::string&)> callback)
    : resolver_(ioc),
      socket_(ioc),
      callback_(std::move(callback)) {
    req_.version(11);
    req_.method(http::verb::get);
    req_.target(target);
    req_.set(http::field::host, host + ":" + port);
    req_.set(http::field::user_agent, BOOST_BEAST_VERSION_STRING);

    resolver_.async_resolve(
        host, port, std::bind(&AsyncHttpClient::on_resolve, shared_from_this(), std::placeholders::_1, std::placeholders::_2));
}

void AsyncHttpClient::start() {
    // The connection process is started in the constructor.
}

void AsyncHttpClient::on_resolve(boost::system::error_code ec, tcp::resolver::results_type results) {
    if (!ec) {
        boost::asio::async_connect(socket_, results,
                                   std::bind(&AsyncHttpClient::on_connect, shared_from_this(), std::placeholders::_1));
    }
}

void AsyncHttpClient::on_connect(boost::system::error_code ec) {
    if (!ec) {
        http::async_write(socket_, req_,
                          std::bind(&AsyncHttpClient::on_write, shared_from_this(), std::placeholders::_1, std::placeholders::_2));
    }
}

void AsyncHttpClient::on_write(boost::system::error_code ec, std::size_t bytes_transferred) {
    boost::ignore_unused(bytes_transferred);

    if (!ec) {
        http::async_read(socket_, buffer_, res_,
                         std::bind(&AsyncHttpClient::on_read, shared_from_this(), std::placeholders::_1, std::placeholders::_2));
    }
}

void AsyncHttpClient::on_read(boost::system::error_code ec, std::size_t bytes_transferred) {
    boost::ignore_unused(bytes_transferred);

    if (!ec) {
        // Call the provided callback with the response body.
        callback_(res_.body());

        // Close the socket.
        boost::system::error_code ignored_ec;
        socket_.shutdown(tcp::socket::shutdown_both, ignored_ec);
    }
}

