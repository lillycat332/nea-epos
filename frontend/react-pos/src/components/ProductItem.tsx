import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import { PropTypes, Typography } from '@mui/material';

interface ProductProps {
	name: string;
}

export default function ProductCard(props : ProductProps) {
  return (
    <Card sx={{ minWidth: 275 }}>
      <CardContent>
				<Typography gutterBottom variant="h5" component="h2">
					{props.name}
				</Typography>
      </CardContent>
      <CardActions>
        <Button size="small">Add to Cart</Button>
      </CardActions>
    </Card>
  );
}