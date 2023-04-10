#include <cpprest/http_listener.h>
#include <cpprest/http_client.h>
#include <cpprest/json.h>
#include <condition_variable>
#include <mutex>

using namespace web;
using namespace http;
using namespace utility;
using namespace http::experimental::listener;
using namespace http::client;

std::mutex request_mutex;
std::condition_variable request_cv;
int max_connections = 300;
int active_connections = 0;

pplx::task<void> handle_request(http_request request) {
    if (request.method() == methods::GET) {
        std::unique_lock<std::mutex> lock(request_mutex);
        request_cv.wait(lock, [] { return active_connections < max_connections; });
        ++active_connections;
        lock.unlock();

        http_client client(U("http://localhost:3000/v1/customer/get"));
        return client.request(methods::GET)
            .then([](http_response response) {
                return response.extract_json();
            })
            .then([&request](json::value json_response) {
                request.reply(status_codes::OK, json_response);
            })
            .then([] {
                std::unique_lock<std::mutex> lock(request_mutex);
                --active_connections;
                request_cv.notify_one();
            });
    } else {
        request.reply(status_codes::MethodNotAllowed, U("Only GET method is allowed."));
        return pplx::task_from_result();
    }
}

int main() {
    uri_builder uri(U("http://localhost:8080"));
    uri.append_path(U("/v1/client/get"));

    http_listener listener(uri.to_uri());
    listener.support([](http_request request) {
        handle_request(request);
    });

    try {
        listener
            .open()
            .then([&listener]() { std::cout << "Listening on " << listener.uri().to_string() << std::endl; })
            .wait();

        std::string line;
        std::getline(std::cin, line);
    } catch (std::exception const &e) {
        std::cout << "Error: " << e.what() << std::endl;
    }

    return 0;
}

