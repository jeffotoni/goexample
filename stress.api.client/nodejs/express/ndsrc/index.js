const express = require("express");
const http = require("http");
const app = express();

const port = process.env.PORT || 8080;

app.get("/ping", (req, res) => {
    res.setHeader("Content-Type", "application/json");
    res.status(200).json({ name: "pong" });
});

app.get("/v1/client/get", (req, res) => {
    res.setHeader("Content-Type", "application/json");

    let rpromisse = new Promise((resolve, reject) => {
        const options = {
            method: 'GET',
            host: '127.0.0.1',
            port: 3000,
            path: '/v1/customer/get',
            headers: {
                'Content-Type': 'application/json; charset=UTF-8',
                'Accept': 'application/json',
                'ID': '0001'
            }
        };
        let request = http.request(options, (res) => {
            if (res.statusCode !== 200) {
                console.error(`Did not get an OK from the server. Code: ${res.statusCode}`);
                res.resume();
                return;
            }

            let data = '';
            res.on('data', (chunk) => {
                data += chunk;
            });

            res.on('close', () => {
                // console.log('Updated data');
                // console.log(JSON.parse(data));
                try {
                    let obj = JSON.parse(data);
                    // console.log('rest::', obj);
                    resolve({
                        statusCode: res.statusCode,
                        data: obj
                    });
                } catch (err) {
                    console.error('rest::end', err);
                    reject(err);
                }
            });
        });
        request.end();
        request.on('error', (err) => {
            console.error(`Encountered an error trying to make a request: ${err.message}`);
        });
    });

    rpromisse.then(({ statusCode, data }) => {
        res.status(200).json(data);
    }, (error) => {
        next(error);
    });
});

app.listen(port);
console.log("Run Server " + port);
