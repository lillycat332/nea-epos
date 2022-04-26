import * as React from "react";
import * as ReactDOMClient from "react-dom/client";
import './index.css';
import AddToCartFAB from "./components/FAB.tsx";
import ResponsiveDrawer from "./components/navdrawer.tsx";

export default function App() {
  return (
    <div>
      <div className="padding-right">
        <ResponsiveDrawer />
      </div>
    </div>
  );
}

const container = document.getElementById('root');
const root = ReactDOMClient.createRoot(container);

root.render(
  <App />,
);