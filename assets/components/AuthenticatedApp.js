import React, { useState, useEffect } from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
  Link,
} from "react-router-dom";
import {
  AppBar,
  Toolbar,
  IconButton,
  Typography,
  makeStyles,
} from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import Home from "./Home";
import CreateDungeon from "./CreateDungeon";
import Dungeon from "./Dungeon";

const useStyles = makeStyles((theme) => ({
  appBar: {
    marginBottom: theme.spacing(4),
  },
}));

export default function AuthenticatedApp() {
  const classes = useStyles();
  return (
    <Router>
      <div>
        <AppBar position="static" className={classes.appBar}>
          <Toolbar>
            <IconButton edge="start" color="inherit" aria-label="menu">
              <MenuIcon />
            </IconButton>
            <Typography variant="h6">Dungeon Crawler</Typography>
          </Toolbar>
        </AppBar>
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/create-dungeon" exact component={CreateDungeon} />
          <Route path="/dungeon/:id" exact component={Dungeon} />
          <Redirect to="/" />
        </Switch>
      </div>
    </Router>
  );
}
