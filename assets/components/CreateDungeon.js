import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import {
  Container,
  Typography,
  TextField,
  Button,
  Grid,
  makeStyles,
} from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  paper: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  heading: {
    marginBottom: theme.spacing(2),
  },
}));

export default function CreateDungeon() {
  const [id, setId] = useState(null);
  const classes = useStyles();
  const handleSubmit = (e) => {
    e.preventDefault();
    fetch("/api/create-dungeon", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        name: e.target.name.value,
        height: e.target.height.value,
        width: e.target.width.value,
      }),
    }).then((response) => {
      if (response.status == 200) {
        response.json().then((data) => {
          setId(data.InsertedID);
        });
      } else {
        console.log("failure");
      }
    });
  };

  return id ? (
    <Redirect to={"/dungeon/" + id} />
  ) : (
    <Container maxWidth="sm" className={classes.paper}>
      <Typography component="h1" variant="h4" className={classes.heading}>
        Create a new Dungeon
      </Typography>
      <form onSubmit={handleSubmit}>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <TextField
              name="name"
              id="name"
              label="Name"
              variant="outlined"
              required
              fullWidth
            ></TextField>
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              name="height"
              id="height"
              label="Height"
              variant="outlined"
              required
              type="number"
              fullWidth
            ></TextField>
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              name="width"
              id="width"
              label="Width"
              variant="outlined"
              required
              type="number"
              fullWidth
            ></TextField>
          </Grid>
          <Grid item xs={12}>
            <Button type="submit" variant="contained" color="primary" fullWidth>
              Create
            </Button>
          </Grid>
        </Grid>
      </form>
    </Container>
  );
}
