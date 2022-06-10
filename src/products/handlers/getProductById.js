import { findProductById } from "../service.js";

export const handler = async (event) => {
  const productId = event.pathParameters?.productId;
  const product = await findProductById(productId);

  return {
    statusCode: 200,
    body: JSON.stringify({
      product,
      event,
    }),
  };
};
