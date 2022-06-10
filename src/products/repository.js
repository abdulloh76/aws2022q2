import { products } from "../mocks/products.js";

export const getAll = () => {
  return products;
};

export const findById = (id) => {
  return products.find((product) => product.id === id);
};
