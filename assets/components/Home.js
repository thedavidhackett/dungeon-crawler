import React, { useEffect, useState } from "react";
import { Container } from "reactstrap";
import { Link } from "react-router-dom";

export default function Home() {
  const [dungeons, setDungeons] = useState([]);
  useEffect(() => {
    fetch("/api/get-dungeons", {
      method: "GET",
      credentials: "include",
    }).then((response) => {
      response.json().then((data) => setDungeons(data));
    });
  }, []);

  return (
    <Container>
      <Link to="/create-dungeon">Create a new dungeon</Link>
      <ul>
        {dungeons.map((dungeon) => {
          return (
            <li key={dungeon.id}>
              <Link to={"/dungeon/" + dungeon.id}>{dungeon.name}</Link>
            </li>
          );
        })}
      </ul>
    </Container>
  );
}
