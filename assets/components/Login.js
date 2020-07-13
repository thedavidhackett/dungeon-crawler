import React from "react";
import { Container, Button, Typography, SvgIcon } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import GoogleIcon from "../svg/google_icon.svg";

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  heading: {
    marginBottom: theme.spacing(4),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function Login() {
  const classes = useStyles();

  return (
    <Container maxWidth="sm">
      <div className={classes.paper}>
        <Typography component="h1" variant="h4" className={classes.heading}>
          Welcome to Dungeon Crawler
        </Typography>
        <Button
          variant="outlined"
          color="primary"
          href="http://localhost:8000/auth/google/login"
          fullWidth
        >
          <SvgIcon component={GoogleIcon} viewBox="0 0 600 476.6" />
          Sign in with Google
        </Button>
      </div>
    </Container>
  );
}
