import React from 'react';
import './App.css';
import { Button, Table } from 'react-bootstrap';

const listUsers = () => {
  console.log("fetch api")
  fetch("/users")
    .then((response) => response.json())
    .then((json) => {
      console.log(json);
    });
}

function App() {
  return (
    <div className="App">
      <Button onClick={listUsers}>Click to show users</Button>
      <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Email</th>
          </tr>
        </thead>
        <tbody>
        </tbody>
      </Table>
    </div>
  );
}

export default App;
