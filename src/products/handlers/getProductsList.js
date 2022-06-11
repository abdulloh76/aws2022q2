import { getAllProducts } from "../service.js";

export const handler = async (event) => {
  const products = await getAllProducts();

  return {
    statusCode: 200,
    headers: {
      "Access-Control-Allow-Origin" : "*",
      "Access-Control-Allow-Credentials" : true
    },
    body: JSON.stringify(products),
  };
};
