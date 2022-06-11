import { getAll, findById } from "./repository.js";

export const getAllProducts = () => {
  return Promise.resolve(getAll());
};

export const findProductsById = (id) => {
  return Promise.resolve(findById(id));
};
