import * as React from "react";
import * as ReactDOMClient from "react-dom/client";
import './index.css';
// @ts-ignore
import ResponsiveDrawer from "./components/navdrawer.tsx";
import 'whatwg-fetch';
import { useState } from "react";

export default function App() {
  return (
    <ResponsiveDrawer />
  );
}

const container = document.getElementById('root');
const root = ReactDOMClient.createRoot(container);

root.render(
  <App />,
);