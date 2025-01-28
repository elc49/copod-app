"use client";

import { Box, Container, Flex, Heading, SimpleGrid, Text } from "@chakra-ui/react";
import { GlobeEarth, Landmark, Lock } from "@/icons/icons";
import { EmailForm } from "@/components/ui/EmailForm";

export default function Home() {
  return (
    <Box bgGradient="to-b" gradientFrom="gray.900" gradientTo="gray.800" h="dvh" color="white">
      <Container mx="auto" px="4" pt="20" pb="32">
        <Box maxW="4xl" mx="auto" textAlign="center">
          <Heading size={{ base: "5xl", md: "6xl", lg: "7xl" }} mb="6" fontWeight="bold">
            Find. Request. Acquire Land Usage Rights
          </Heading>
          <Text fontSize="xl" color="gray.300" mb="12">
            Join the first real-world land ownership platform on Ethereum blockchain. Secure, Transparent, and Borderless land use agreements
          </Text>
          <Box maxW="md" mx="auto">
            <EmailForm />
          </Box>
        </Box>
      </Container>
      <Box py="20" bg="gray.800">
        <Container mx="auto" px="4">
          <SimpleGrid columns={[1, 3]} maxW="5xl" mx="auto" gap="12">
            <Box textAlign="center">
              <Flex justify="center" mb="4">
                <Lock boxSize="48px" color="blue.500" />
              </Flex>
              <Heading size="xl" fontWeight="bold" mb="2">
                Secure & Trustless Agreements
              </Heading>
              <Text color="gray.400">
                Smart contracts ensure safe and transparent land usage agreements
              </Text>
            </Box>
            <Box textAlign="center">
              <Flex justify="center" mb="4">
                <Landmark boxSize="48px" color="blue.500" />
              </Flex>
              <Heading size="xl" fontWeight="bold" mb="2">
                Verified land and land owners
              </Heading>
              <Text color="gray.400">
                Government issued titles are zkp-verified and on the blockchain
              </Text>
            </Box>
            <Box textAlign="center">
              <Flex justify="center" mb="4">
                <GlobeEarth boxSize="48px" color="blue.500" />
              </Flex>
              <Heading size="xl" fontWeight="bold" mb="2">
                Arable. Semi-arid. Arid
              </Heading>
              <Text color="gray.400">
                Farming, Camping. Find your use case
              </Text>
            </Box>
          </SimpleGrid>
        </Container>
      </Box>
    </Box>
  );
}
