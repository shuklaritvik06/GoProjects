const formElement = document.querySelector("#form");
const decompressElement = document.querySelector("#decompress");
const file = document.querySelector("#file");
const decompress = document.querySelector("#decompress_file");
const formData = new FormData();

const handleSubmit = (event) => {
  event.preventDefault();

  for (const myfile of file.files) {
    formData.append("file", myfile);
  }

  fetch("http://localhost:5000/compress", {
    method: "post",
    headers: {
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFua2l0amFpbiIsIkZpcnN0X05hbWUiOiJBbmtpdCIsIkxhc3RfTmFtZSI6IkphaW4iLCJleHAiOjE2ODAwOTYxNzR9.BQ2iL8qKQ7c5dGC6GxFs7YBSeN-8Zryvj7s2pbp9yak"
    },
    body: formData
  })
    .then((res) => res.arrayBuffer())
    .then((data) => {
      saveByteArray([data], "compressed.zip");
    })
    .catch((error) => ("Something went wrong!", error));
};

const handleDecompress = (event) => {
  console.log(event);
  event.preventDefault();

  for (const myfile of decompress.files) {
    formData.append("file", myfile);
  }

  fetch("http://localhost:5000/decompress", {
    method: "post",
    headers: {
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFua2l0amFpbiIsIkZpcnN0X05hbWUiOiJBbmtpdCIsIkxhc3RfTmFtZSI6IkphaW4iLCJleHAiOjE2ODAwOTYxNzR9.BQ2iL8qKQ7c5dGC6GxFs7YBSeN-8Zryvj7s2pbp9yak"
    },
    body: formData
  })
    .then((res) => res.arrayBuffer())
    .then((data) => {
      saveByteArray([data], "decompressed");
    })
    .catch((error) => ("Something went wrong!", error));
};

const saveByteArray = (function () {
  var a = document.createElement("a");
  document.body.appendChild(a);
  a.style = "display: none";
  return function (data, name) {
    const blob = new Blob(data, { type: "octet/stream" });
    url = window.URL.createObjectURL(blob);
    a.href = url;
    a.download = name;
    a.click();
    window.URL.revokeObjectURL(url);
  };
})();

formElement.addEventListener("submit", handleSubmit);
decompressElement.addEventListener("submit", handleDecompress);
