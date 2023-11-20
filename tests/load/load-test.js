import http from "k6/http";

const API = __ENV.BASE_URL;

export default function () {
  let res = http.get(API);

  res = res.submitForm({
    formSelector: "form",
    fields: { meme: "knock knock" },
  });
}

export const options = {
  stages: [
    { duration: "20s", target: 100 },
    { duration: "20s", target: 100 },
    { duration: "20s", target: 0 },
  ],
};