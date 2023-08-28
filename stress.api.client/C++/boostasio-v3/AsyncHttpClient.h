#ifndef ASYNC_HTTP_CLIENT_H
#define ASYNC_HTTP_CLIENT_H

#include <boost/asio.hpp>
#include <boost/beast.hpp>
#include <functional>
#include <string>

namespace http = boost::beast::http;
using tcp = boost::asio::ip::tcp;

class AsyncHttpClient : public std::enable_shared_from_this<AsyncHttpClient> {
public:
    AsyncHttpClient(boost::asio::io_context& ioc, const std::string& host, const std::string& port,
                    const std::string& target, std::function<void(const std::string&)> callback);

    void start();

private:
    void on_resolve(boost::system::error_code ec, tcp::resolver::results_type results);
    void on_connect(boost::system::error_code ec);
    void on_write(boost::system::error_code ec, std::size_t bytes_transferred);
    void on_read(boost::system::error_code ec, std::size_t bytes_transferred);

    tcp::resolver resolver_;
    tcp::socket socket_;
    http::request<http::empty_body> req_;
    http::response<http::string_body> res_;
    boost::beast::flat_buffer buffer_;
    std::function<void(const std::string&)> callback_;
};

#endif // ASYNC_HTTP_CLIENT_H
