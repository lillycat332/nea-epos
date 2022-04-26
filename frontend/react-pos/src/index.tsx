import * as React from "react";
import * as ReactDOMClient from "react-dom/client";
import './index.css';
import ResponsiveDrawer from "./components/navdrawer.tsx";
import ProductCard from './components/ProductItem.tsx';

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