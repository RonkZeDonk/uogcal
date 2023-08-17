/*
  You can build the bookmarklet yourself using other websites or just
  encoding this file as a URI component and sticking it in the middle of the
  following:

    let bookmarklet = `javascript:(()=>{{ ${JS_STRING} }})()`

  The latter is the method used by this website.

  TODO make sure that the user is registered in the course before adding it
    * Pre semester message:    "Registered, but not started"
    * During semester message: ???
*/

let list = document.getElementsByClassName("schedule-list")[0].children;

let termStr = document.getElementById("schedule-activeterm-text").innerText;
let term = termStr.replace(/(.)(?:.*)(..)/g, "$1$2");

let data = {
  term,
  courses: [],
};
for (let li of list) {
  let title = li.getElementsByClassName("schedule-listitem-header-title")[0]
    .innerText;

  data.courses.push({ title });
}

let a = window.document.createElement("a");
a.href = window.URL.createObjectURL(
  new Blob([JSON.stringify(data)], { type: "application/json" })
);
a.download = "courses.json";

document.body.appendChild(a);
a.click();
document.body.removeChild(a);
