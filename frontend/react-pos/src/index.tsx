import * as React from "react";
import * as ReactDOMClient from "react-dom/client";
import './index.css';
import ResponsiveDrawer from "./components/navdrawer.tsx";
import ProductCard from './components/ProductItem.tsx';

const products = [['Banana', 1.0], ['Passionfruit', 2.0], ['Dragonfruit', 3.0], ['Strawbebby', 4.0], ['Starfruit', 5.0]];
const listItems = products.map((product) =>
  <ProductCard name={product[0]} price={product[1]} imagePath="logo512.png"/>
);

export default function App() {
  return (
    <div className="padding-right">
        <ResponsiveDrawer />
    </div>
  );
}

const container = document.getElementById('root');
const root = ReactDOMClient.createRoot(container);

root.render(
  <App />,
);