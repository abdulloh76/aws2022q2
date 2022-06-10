import { getAll, findById } from "./repository.js";

export const getAllProducts = () => {
  return Promise.resolve(getAll());
};

export const findProductById = (id) => {
  return Promise.resolve(findById(id));
};
