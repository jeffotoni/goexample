#include <cpprest/http_listener.h>
#include <cpprest/http_client.h>
#include <cpprest/json.h>

using namespace web;
using namespace http;
using namespace utility;
using namespace http::experimental::listener;
using namespace http::client;

// sudo apt-get install libcpprest-dev
// g++ main.cpp -o cppserver -lcpprest -lboost_system -lcrypto -lssl -lpthread
void handle_request(http_request request) {
    if (request.method() == methods::GET) {
        http_client client(U("http://localhost:3000/v1/customer/get"));
        client.request(methods::GET)
            .then([](http_response response) {
                return response.extract_json();
            })
            .then([&request](json::value json_response) {
                request.reply(status_codes::OK, json_response);
            })
            .wait();
    } else {
        request.reply(status_codes::MethodNotAllowed, U("Only GET method is allowed."));
    }
}

int main() {
    uri_builder uri(U("http://localhost:8080"));
    uri.append_path(U("/v1/client/get"));

    http_listener listener(uri.to_uri());
    listener.support(handle_request);

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

