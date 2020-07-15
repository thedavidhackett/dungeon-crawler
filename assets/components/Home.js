import React, { useEffect, useState } from "react";
import { Container, Row, Col, Button } from "reactstrap";
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
      <Row>
        <Col xs={12} md={{ size: 6, offset: 3 }}>
          <div className="dungeons">
            <h4>Your Dungeons</h4>
            {dungeons.length > 0 ? (
              <ul className="list-group dungeons-list">
                {dungeons.map((dungeon) => {
                  return (
                    <li key={dungeon.id} className="list-group-item">
                      <Link to={"/dungeon/" + dungeon.id}>{dungeon.name}</Link>
                    </li>
                  );
                })}
              </ul>
            ) : (
              <p>No Dungeons Yet</p>
            )}
            <Link to="/create-dungeon">
              <Button>Create a new dungeon</Button>
            </Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
}
