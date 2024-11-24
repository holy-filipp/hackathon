import {Flex, Box, Text, Button, CheckboxGroup } from "@radix-ui/themes";
import './header.css'
import icon1 from './h_img/dis.svg'
import icon2 from './h_img/prog.svg'
import icon3 from './h_img/place.svg'
import icon4 from './h_img/part.svg'
import icon5 from './h_img/sex.svg'
import icon6 from './h_img/age.svg'
import icon7 from './h_img/dl.svg'
import icon8 from './h_img/type.svg'
import icon9 from './h_img/stage.svg'
import icon10 from './h_img/compos.svg'
import icon11 from './h_img/level.svg'

export default function Header() {
    function a() {
        const elements = document.querySelectorAll('.sbut');
        elements.forEach(element => {
          if (element.style.display === 'none') {
            element.style.display = 'block';
          } else {
            element.style.display = 'none';
          }
        });
      }
      
   
	return (
        
		<Flex direction='column' className="header">
            <Box className="Box"><Text className="logo">Glory team</Text></Box>
            <Flex direction='column' gapY='7px'>
                <Button className="br"  onClick={a}>Вид спорта</Button>
                <CheckboxGroup.Root defaultValue={['0']} name="example" className="sbut">
                  <CheckboxGroup.Item value="1">баскетбол</CheckboxGroup.Item>
                  <CheckboxGroup.Item value="2">бейсбол</CheckboxGroup.Item>
                  <CheckboxGroup.Item value="3">волейбол</CheckboxGroup.Item>
                </CheckboxGroup.Root>
                <Button className="br"><img src={icon1}/>дисциплина</Button>
                <Button className="br"><img src={icon2} />программа</Button>
                <Button className="br"><img src={icon3} />место проведения</Button>
                <Button className="br"><img src={icon4}/>количество участников</Button>
                <Button className="br"><img src={icon5}/>пол</Button>
                <Button className="br"><img src={icon6}/>возрастная группа</Button>
                <Button className="br"><img src={icon7}/>сроки проведения</Button>
                <Button className="br"><img src={icon8}/>тип соревнования</Button>
                <Button className="br"><img src={icon9}/>этап</Button>
                <Button className="br"><img src={icon10}/>состав</Button>
                <Button className="br"><img src={icon11}/>уровень</Button>
            </Flex>

        </Flex>

	);
}