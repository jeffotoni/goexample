const { Worker, isMainThread, parentPort } = require('worker_threads');
const os = require('os');

if (isMainThread) {
    const express = require('express');
    const app = express();
    const port = process.env.PORT || 8080;

    app.get('/ping', (req, res) => {
        res.setHeader('Content-Type', 'application/json');
        res.status(200).json({ name: 'pong' });
    });

    app.get('/v1/client/get', (req, res) => {
        const worker = new Worker(__filename);

        worker.on('message', ({ statusCode, data }) => {
            res.status(statusCode).json(data);
        });

        worker.on('error', (error) => {
            res.status(500).json({ error: 'Internal server error' });
        });

        worker.on('exit', (code) => {
            if (code !== 0) {
                res.status(500).json({ error: 'Internal server error' });
            }
        });
    });

    app.listen(port, () => {
        console.log(`Server is running on port ${port}`);
    });
} else {
    const http = require('http');

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

    const request = http.request(options, (res) => {
        if (res.statusCode !== 200) {
            console.error(`Did not get an OK from the server. Code: ${res.statusCode}`);
            res.resume();
            return;
        }

        let data = '';
        res.on('data', (chunk) => {
            data += chunk;
        });

        res.on('end', () => {
            try {
                const obj = JSON.parse(data);
                parentPort.postMessage({
                    statusCode: res.statusCode,
                    data: obj
                });
            } catch (err) {
                console.error('Error:', err);
            }
        });
    });

    request.end();
    request.on('error', (err) => {
        console.error(`Encountered an error trying to make a request: ${err.message}`);
    });
}