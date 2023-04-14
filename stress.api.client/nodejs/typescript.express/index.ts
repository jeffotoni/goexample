import * as express from "express";
import { Request, Response, NextFunction } from "express";
import * as http from "http";
import { RequestOptions } from "http";

const app = express();
const port = process.env.PORT || 8080;

app.get("/ping", (req: Request, res: Response) => {
  res.setHeader("Content-Type", "application/json");
  res.status(200).json({ name: "pong" });
});

app.get("/v1/client/get", async (req: Request, res: Response, next: NextFunction) => {
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
    res.status(statusCode).json(data);
  } catch (error) {
    next(error);
  }
});

app.listen(port, () => {
  console.log("Run Server " + port);
});
