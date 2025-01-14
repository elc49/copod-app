"use client";

import { Box, Container, Flex, Heading, SimpleGrid, Text } from "@chakra-ui/react";
import { GlobeEarth, Landmark, Lock } from "@/icons/icons";
import { EmailForm } from "@/components/ui/EmailForm";

export default function Home() {
  return (
    <Box bgGradient="to-b" gradientFrom="gray.900" gradientTo="gray.800" h="dvh" color="white">
      <Container mx="auto" px="4" pt="20" pb="32">
        <Box maxW="4xl" mx="auto" textAlign="center">
          <Heading size="5xl" mb="6" fontWeight="bold">
            The Future of Land Ownership is Digital
          </Heading>
          <Text fontSize="xl" color="gray.300" mb="12">
            Join the first decentralized land ownership platform on Ethereum. Secure, transparent, and borderless property transactions
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
                Secure Transactions
              </Heading>
              <Text color="gray.400">
                Smart contracts ensure safe and transparent property transactions
              </Text>
            </Box>
            <Box textAlign="center">
              <Flex justify="center" mb="4">
                <Landmark boxSize="48px" color="blue.500" />
              </Flex>
              <Heading size="xl" fontWeight="bold" mb="2">
                Verified properties
              </Heading>
              <Text color="gray.400">
                All listings are verified and authenticated on the blockchain
              </Text>
            </Box>
            <Box textAlign="center">
              <Flex justify="center" mb="4">
                <GlobeEarth boxSize="48px" color="blue.500" />
              </Flex>
              <Heading size="xl" fontWeight="bold" mb="2">
                Global access
              </Heading>
              <Text color="gray.400">
                Buy property rights from anywhere in the world
              </Text>
            </Box>
          </SimpleGrid>
        </Container>
      </Box>
    </Box>
  );
}
