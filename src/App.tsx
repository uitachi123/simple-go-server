import React from 'react';
import './App.css';
import { Button } from 'react-bootstrap';

const listUsers = () => {
  console.log("fetch api")
  fetch("/users")
    .then((response) => console.log(response))
  // .then((json) => {
  //   console.log(json);
  // });
}

function App() {
  return (
    <div className="App">
      <Button onClick={listUsers}>Click to show users</Button>
    </div>
  );
}

export default App;
