import * as http from "http";
import { RequestOptions } from "http";
import { IncomingMessage, ServerResponse } from "http";

const port = process.env.PORT || 8080;

const server = http.createServer(async (req: IncomingMessage, res: ServerResponse) => {
  if (req.url === "/ping" && req.method === "GET") {
    res.setHeader("Content-Type", "application/json");
    res.statusCode = 200;
    res.end(JSON.stringify({ name: "pong" }));
  } else if (req.url === "/v1/client/get" && req.method === "GET") {
    res.setHeader("Content-Type", "application/json");

    try {
      const options: RequestOptions = {
        method: "GET",
        host: "127.0.0.1",
        port: 3000,
        path: "/v1/customer/get",
        headers: {
          "Content-Type": "application/json; charset=UTF-8",
          Accept: "application/json",
          ID: "0001",
        },
      };

      const fetchData = (): Promise<{ statusCode: number; data: any }> =>
        new Promise((resolve, reject) => {
          const request = http.request(options, (response) => {
            if (response.statusCode !== 200) {
              console.error(`Did not get an OK from the server. Code: ${response.statusCode}`);
              response.resume();
              return;
            }

            let data = "";
            response.on("data", (chunk) => {
              data += chunk;
            });

            response.on("end", () => {
              try {
                const obj = JSON.parse(data);
                resolve({
                  statusCode: response.statusCode,
                  data: obj,
                });
              } catch (err) {
                console.error("rest::end", err);
                reject(err);
              }
            });
          });

          request.end();
          request.on("error", (err) => {
            console.error(`Encountered an error trying to make a request: ${err.message}`);
          });
        });

      const { statusCode, data } = await fetchData();
      res.statusCode = statusCode;
      res.end(JSON.stringify(data));
    } catch (error) {
      res.statusCode = 500;
      res.end("Internal Server Error");
    }
  } else {
    res.statusCode = 404;
    res.end("Not Found");
  }
});

server.listen(port, () => {
  console.log("Run Server nodejs:" + port);
});
