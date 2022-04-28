import * as React from "react";
import * as ReactDOMClient from "react-dom/client";
import './index.css';
// @ts-ignore
import MainView from "./components/MainView.tsx";
import 'whatwg-fetch';

export default function App() {
  return (
    <MainView />
  );
}

const container = document.getElementById('root');
const root = ReactDOMClient.createRoot(container);

root.render(
  <App />,
);