import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import { Typography } from '@mui/material';
import CardMedia from '@mui/material/CardMedia';

interface ProductProps {
  name: string;
  imagePath: string;
  price: number;
  currency: string;
  addToCart: Function;
}

export default function ProductCard(props: ProductProps) {
  return (
    <Card sx={{ minWidth: 275 }}>
      <CardMedia
        component="img"
        height="140"
        image={props.imagePath}
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="h2">
          {props.name}
        </Typography>
        <Typography variant="body2" color="textSecondary" component="p">
          {props.currency}{props.price}
        </Typography>
      </CardContent>
      <CardActions>
        <Button size="small" onClick={() => this.addToCart}>Add to Cart</Button>
      </CardActions>
    </Card>
  );
}