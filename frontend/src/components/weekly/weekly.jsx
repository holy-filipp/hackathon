import {Grid, Box, Text, Button, Heading, Flex } from "@radix-ui/themes";
import Calendar from "../calendar/calendar.jsx";
import img from "../calendar/img/line.svg";

export default function Weekly() {
	return (
        <Grid columns="2" className="ml5">
        <Grid columns="7">
            <div className="item">
            <Flex className="item_top">
                <p>fff</p>
                <p>fff</p>
            </Flex>
            <Flex className="item_bottom">
                <Box className="st_item_bottom">
                <p>ff</p>
                </Box>
            </Flex>
            </div>
        </Grid>

        <Grid columns="1" className="weekly">
            <Heading className="info">
            <p>НА ЭТОЙ НЕДЕЛЕ</p>
            <img src={img} alt="" width="80%" height="3px" />
            </Heading>

            <Box className="point">
            <p>{"я сосу хуй"}</p>
            <p>{"Нет я сосоу"}</p>
            </Box>
        </Grid>
        </Grid>
	);
}