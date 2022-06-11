import { findProductsById } from "../service.js";

export const handler = async (event) => {
  const { productId = "" } = event.pathParameters;
  const product = await findProductsById(productId);

  if (product) {
    return {
      statusCode: 200,
      headers: {
        "Access-Control-Allow-Origin" : "*"
      },
      body: JSON.stringify(product),
    };
  } else {
    return {
      statusCode: 404,
      headers: {
        "Access-Control-Allow-Origin" : "*",
        "Access-Control-Allow-Credentials" : true
      },
      body: JSON.stringify({
        error: { message: "product with id not found" }
      }),
    };
  }
};
