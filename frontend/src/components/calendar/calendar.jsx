import {
  Grid,
  Box,
  Text,
  Button,
  Heading,
  TextField,
  TabNav,
  Flex,
  Dialog
} from "@radix-ui/themes";
import "./calendar.css";
import img from "./img/line.svg";

export default function Calendar() {
  const month = ["Январь", "Февраль"];
  return (
    <Grid columns="2"  gapx='2' className="ml5 guru">
      <Grid columns="7">
        <Dialog.Root>
          <Dialog.Trigger>
            <Button className="item">
              <div className="flex">
                <p>fff</p>
                <p>fff</p>
              </div>
                <p className="name">ff</p>
            </Button>
          </Dialog.Trigger>

          <Dialog.Content maxWidth="450px" maxHeight="700px">
            <Dialog.Title align='center'>
              {22222}
            </Dialog.Title>
            <Dialog.Title className="dialog_item">
              <p>sdf</p>
              <p>saf</p>
            </Dialog.Title>
          </Dialog.Content>
        </Dialog.Root>
        
        
      </Grid>

      <Grid columns="1" className="weekly">
        <Heading className="info">
          <p>НА ЭТОЙ НЕДЕЛЕ</p>
          <img src={img} alt="" width="80%" height="3px" />
        </Heading>
        <Box className="point">
          <p>{"я сосу"}</p>
          <p>{"Нет я сосоу"}</p>
        </Box>
      </Grid>
    </Grid>
  );
}
