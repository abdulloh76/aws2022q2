import { findProductById, getAllProducts } from "./service.js";

const getProducts = async (event) => {
  const products = await getAllProducts();

  return {
    statusCode: 200,
    body: JSON.stringify({
      products,
      event,
    }),
  };
};

const getProductById = async (event) => {
  const product = await findProductById();

  return {
    statusCode: 200,
    body: JSON.stringify({
      product,
      event,
    }),
  };
};

export { getProducts, getProductById };
