import React from 'react';
import {Container} from "semantic-ui-react";
import Supplier from "./components/supplierForm";
import './App.css';

function App() {
  return (
    <div className="App">
      <Container>
        <Supplier/>
      </Container>
    </div>
  );
}

export default App;
