// npm init -y
// npm install fastify
// npm install node-fetch
// npm install @types/node-fetch --save-dev

const fastify = require("fastify");
const app = fastify({ logger: false });
const port = process.env.PORT || 8080;

app.get("/ping", async(_, reply) => {
    reply.type("application/json").code(200);
    return { name: "pong" };
});

app.get("/v1/client/get", async(_, reply) => {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json; charset=UTF-8",
            Accept: "application/json",
            ID: "0001",
        },
    };

    try {
        const fetch = await
        import ("node-fetch");
        const response = await fetch.default("http://127.0.0.1:3000/v1/customer/get", requestOptions);
        if (response.status !== 200) {
            throw new Error(`Did not get an OK from the server. Code: ${response.status}`);
        }

        const data = await response.json();
        reply.type("application/json").code(200);
        return data;
    } catch (error) {
        //app.log.error(error);
        reply.type("application/json").code(500);
        return { error: "Internal Server Error" };
    }
});

app.listen({ port: port }, (err, address) => {
    if (err) {
        app.log.error(err);
        process.exit(1);
    }
    // app.log.info(`Server listening at ${address}`);
});