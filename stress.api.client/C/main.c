#include <aco.h>
#include <curl/curl.h>
#include <ev.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#define PORT 8080
#define BUFFER_SIZE 1024

// Callback for handling libcurl response data
static size_t write_data(void* ptr, size_t size, size_t nmemb, void* stream) {
    size_t written = fwrite(ptr, size, nmemb, (FILE*)stream);
    return written;
}

// gcc main.c -o coro_client -laco -lev -lcurl -lpthread
// Coroutine function to fetch customer data
void fetch_customer_data(aco_t* co) {
    CURL* curl;
    FILE* headerfile;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_DEFAULT);
    curl = curl_easy_init();
    if (curl) {
        headerfile = tmpfile();
        if (headerfile == NULL) {
            perror("fopen");
            exit(EXIT_FAILURE);
        }

        curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:3000/v1/customer/get");
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, headerfile);

        res = curl_easy_perform(curl);
        if (res != CURLE_OK) {
            fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
        } else {
            printf("HTTP request completed successfully\n");
        }

        curl_easy_cleanup(curl);
        fclose(headerfile);
    }

    curl_global_cleanup();
}

// Coroutine entry function
void aco_entry(aco_t* co, void* arg) {
    fetch_customer_data(co);
}

int main() {
    // Initialize coroutine
    aco_t* main_co = aco_create(NULL, NULL, 0, NULL, NULL);
    aco_t* co = aco_create(main_co, NULL, 0, aco_entry, NULL);

    // Run the coroutine
    aco_resume(co);

    // Clean up
    aco_destroy(co);
    aco_destroy(main_co);
    return 0;
}

