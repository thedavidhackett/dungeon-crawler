import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import "./index.css";

function App() {
  const [user, setUser] = useState(null);
  useEffect(() => {
    fetch("/auth/check-logged-in", {
      method: "GET",
      credentials: "include",
    }).then((response) => {
      response.json().then((data) => setUser(data));
    });
  }, []);

  return user ? (
    <div>Hello World</div>
  ) : (
    <div>
      Not Logged In{" "}
      <a href="http://localhost:8000/auth/google/login">Log in with google</a>
    </div>
  );
}

ReactDOM.render(<App />, document.getElementById("app"));
