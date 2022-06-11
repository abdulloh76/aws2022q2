import { getAll } from "../../src/products/repository";

describe("getProductsRepository", () => {
  it("should return array of products", () => {
    const products = getAll();
    expect(Array.isArray(products)).toBeTruthy();
  });
});
