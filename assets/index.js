import "bootstrap/dist/css/bootstrap.min.css";
import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { RecoilRoot, useRecoilState } from "recoil";
import AuthenticatedApp from "./components/AuthenticatedApp";
import Login from "./components/Login";
import { userState } from "./atoms";

function App() {
  const [user, setUser] = useRecoilState(userState);
  useEffect(() => {
    fetch("/auth/check-logged-in", {
      method: "GET",
      credentials: "include",
    }).then((response) => {
      response.json().then((data) => setUser(data));
    });
  }, []);

  return user ? <AuthenticatedApp /> : <Login />;
}

ReactDOM.render(
  <RecoilRoot>
    <App />
  </RecoilRoot>,
  document.getElementById("app")
);
