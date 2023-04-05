"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var express = require("express");
var http = require("http");
var app = express();
var port = process.env.PORT || 8080;
app.get("/ping", function (req, res) {
    res.setHeader("Content-Type", "application/json");
    res.status(200).json({ name: "pong" });
});
app.get("/v1/client/get", function (req, res, next) { return __awaiter(void 0, void 0, void 0, function () {
    var options_1, fetchData, _a, statusCode, data, error_1;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                res.setHeader("Content-Type", "application/json");
                _b.label = 1;
            case 1:
                _b.trys.push([1, 3, , 4]);
                options_1 = {
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
                fetchData = function () {
                    return new Promise(function (resolve, reject) {
                        var request = http.request(options_1, function (response) {
                            if (response.statusCode !== 200) {
                                console.error("Did not get an OK from the server. Code: ".concat(response.statusCode));
                                response.resume();
                                return;
                            }
                            var data = "";
                            response.on("data", function (chunk) {
                                data += chunk;
                            });
                            response.on("end", function () {
                                try {
                                    var obj = JSON.parse(data);
                                    resolve({
                                        statusCode: response.statusCode,
                                        data: obj,
                                    });
                                }
                                catch (err) {
                                    console.error("rest::end", err);
                                    reject(err);
                                }
                            });
                        });
                        request.end();
                        request.on("error", function (err) {
                            console.error("Encountered an error trying to make a request: ".concat(err.message));
                        });
                    });
                };
                return [4 /*yield*/, fetchData()];
            case 2:
                _a = _b.sent(), statusCode = _a.statusCode, data = _a.data;
                res.status(statusCode).json(data);
                return [3 /*break*/, 4];
            case 3:
                error_1 = _b.sent();
                next(error_1);
                return [3 /*break*/, 4];
            case 4: return [2 /*return*/];
        }
    });
}); });
app.listen(port, function () {
    console.log("Run Server " + port);
});
