import { BrowserRouter as  Router} from 'react-router-dom';
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "@radix-ui/themes/styles.css";
import { Theme } from "@radix-ui/themes";
import { Grid } from "@radix-ui/themes";
import Header from "./components/header/header.jsx";
import Main from "./components/center/center.jsx";
import "./main.css";
createRoot(document.getElementById("root")).render(
  <StrictMode>
    <Router>
      <html>
        <body>
          <Theme>
            <Grid columns="2" className="grid">
              <Header className="head"></Header>
              <Main></Main>
            </Grid>
          </Theme>
        </body>
      </html>
    </Router>
  </StrictMode>
);
