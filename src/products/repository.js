import products from "../mocks/products.js";

const getAll = () => {
  return products;
};

const findById = (id) => {
  return products.find((product) => product.id === id);
};

export { getAll, getById };
