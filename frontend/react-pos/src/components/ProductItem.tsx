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
}

const currencyType = "GBP";

function currencyTypeCheck() {
  if (currencyType === 'USD') {
    return '$';
  } else if (currencyType === 'EUR') {
    return '€';
  } else if (currencyType === 'GBP') {
    return '£';
  }
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
          {currencyTypeCheck()}{props.price}
        </Typography>
      </CardContent>
      <CardActions>
        <Button size="small">Add to Cart</Button>
      </CardActions>
    </Card>
  );
}