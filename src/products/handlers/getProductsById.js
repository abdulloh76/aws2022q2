import { findProductsById } from "../service.js";

export const handler = async (event) => {
  const productId = event.pathParameters?.productId;
  const product = await findProductsById(productId);

  if (product) {
    return {
      statusCode: 200,
      body: JSON.stringify({
        product,
        event,
      }),
    };
  } else {
    return {
      statusCode: 404,
      body: JSON.stringify({
        error: { message: "product with id not found" },
        event,
      }),
    };
  }
};
