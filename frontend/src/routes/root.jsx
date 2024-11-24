import { Theme } from "@radix-ui/themes";
import {Grid} from "@radix-ui/themes";
import RootPanel from '../Rootpanel.jsx';
import Header from '../components/header/header.jsx'
import Main from '../components/center/center.jsx'
import '../main.css'

export default function Root() {
    return (
        <html>
        <body>
            <Theme>
                <RootPanel/>
      <Grid columns="2" className='grid'>
        <Header className='head'></Header>
        <Main></Main>
      </Grid>


            </Theme>
        </body>
    </html>
    );
  }