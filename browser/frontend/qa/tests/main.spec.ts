import { test, expect } from '@playwright/test';

test('has html lang attribute set', async ({ page }) => {
    await page.goto('http://localhost:8080/');

    // Expect a title "to contain" a substring.
    const htmlElement = page.locator("html");
    await expect(htmlElement).toHaveAttribute("lang", "en");
});

test('has title with right value', async ({ page }) => {
  await page.goto('http://localhost:8080/');

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/Hery - Playground/);
});

test('has h1 with right value', async ({ page }) => {
  await page.goto('http://localhost:8080/');
  await expect(page.locator("h1")).toHaveText("Hery - Playground 🛝");
});
