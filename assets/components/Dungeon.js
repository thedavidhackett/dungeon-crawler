import React, { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import { useRecoilValue } from "recoil";
import { userState } from "../atoms";
import CharacterForm from "./CharacterForm";
import { Container, Row, Col } from "reactstrap";

export default function Dungeon() {
  const { id } = useParams();
  const user = useRecoilValue(userState);
  const [dungeon, setDungeon] = useState({});
  const [dm, setDm] = useState(false);
  const tileSize = 50;
  const canvasRef = useRef(null);
  const socket = new WebSocket("ws://localhost:8000/api/ws/" + id);
  const [noCharacter, setNoCharacter] = useState(true);
  const [newCharacter, setNewCharacter] = useState(null);

  let connect = (cb) => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = (msg) => {
      console.log(msg);
      cb(msg);
    };

    socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };
  };

  let sendMsg = (msg) => {
    console.log("sending msg: ", msg);
    socket.send(msg);
  };

  useEffect(() => {
    fetch("/api/get-dungeon/" + id, {
      method: "GET",
      credentials: "include",
    }).then((response) => {
      response.json().then((data) => {
        setDungeon(data);
        if (data.userId == user.id) {
          setDm(true);
          setNoCharacter(false);
        } else {
          setDm(false);
          if (data.characters) {
            for (let character of data.character) {
              if (character.userId == user.id) {
                setNoCharacter(false);
                break;
              }
            }
          }
        }
      });
    });
    connect(callBack);
  }, []);

  useEffect(() => {
    drawTileMap(dungeon.height, dungeon.width);
    if (dungeon.characters) {
      for (let character of dungeon.characters) {
        draw(character);
      }
    }
  }, [dungeon]);

  const drawTileMap = (height, width) => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    for (let x = 0; x < width * tileSize; x += tileSize) {
      for (let y = 0; y < height * tileSize; y += tileSize) {
        ctx.strokeRect(x, y, tileSize, tileSize);
      }
    }
  };

  function draw(character) {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");
    let img = new Image();
    img.onload = function () {
      ctx.drawImage(
        img,
        character.positionX - 20,
        character.positionY - 20,
        40,
        40
      );
    };
    img.src = "/images/" + character.image;
  }

  const callBack = (msg) => {
    setDungeon(JSON.parse(msg.data));
  };

  const createNpc = (e) => {
    e.preventDefault();
    setNewCharacter({
      name: e.target.name.value,
      npc: true,
      image: e.target.image.value,
    });
  };

  const createCharacter = (e) => {
    e.preventDefault();
    setNewCharacter({
      name: e.target.name.value,
      npc: false,
      userId: user.id,
      image: e.target.image.value,
    });
  };

  const placeCharacter = (e) => {
    let rect = e.target.getBoundingClientRect();
    let x = e.clientX - rect.left;
    let y = e.clientY - rect.top;
    let newDungeon = { ...dungeon };
    newCharacter.positionX = x;
    newCharacter.positionY = y;
    newDungeon.characters.push({ ...newCharacter });
    setNewCharacter(null);
    sendMsg(JSON.stringify(newDungeon));
  };

  const moveCharacter = (e) => {
    let rect = e.target.getBoundingClientRect();
    let x = e.clientX - rect.left;
    let y = e.clientY - rect.top;
    let newDungeon = { ...dungeon };
    if (newDungeon.characters && newDungeon.characters[0]) {
      newDungeon.characters[0].positionX = Math.floor(x);
      newDungeon.characters[0].positionY = Math.floor(y);
      sendMsg(JSON.stringify(newDungeon));
    }
  };

  return (
    <Container>
      <h4>{dungeon.name}</h4>
      <Row>
        {dm && (
          <Col xs={12}>
            {!newCharacter && (
              <CharacterForm onSubmit={createNpc} label="Create NPC" />
            )}
          </Col>
        )}
        {noCharacter && (
          <Col xs={12}>
            <CharacterForm
              onSubmit={createCharacter}
              label="Create Character"
            />
          </Col>
        )}
        <Col xs={12}>
          <ul>
            {dungeon.characters
              ? dungeon.characters.map((c) => <li>{c.name}</li>)
              : ""}
          </ul>
        </Col>
        <Col item xs={12} className="canvas-container">
          {newCharacter && <label>Place character</label>}
          <canvas
            ref={canvasRef}
            width={(dungeon.width ? dungeon.width : 0) * tileSize}
            height={(dungeon.height ? dungeon.height : 0) * tileSize}
            onClick={newCharacter ? placeCharacter : moveCharacter}
          />
        </Col>
      </Row>
    </Container>
  );
}
