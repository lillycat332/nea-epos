import * as React from "react";
import Box from '@mui/material/Box';
import Fab from '@mui/material/Fab';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';

export default function AddToCartFAB() {
	return (
		<Box sx={{ '& > :not(style)': { m: 1 } }}>
			<Fab color="primary" aria-label="add" className="bottom-left">
        <AddShoppingCartIcon />
      </Fab>
		</Box>
	)
}