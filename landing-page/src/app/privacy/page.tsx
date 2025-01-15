"use client";

import { Box, Container, Heading, Link, Text } from "@chakra-ui/react";

function Page() {
  return (
    <Box mx="auto" h="dvh">
      <Container maxW="3xl" py="8" spaceY="4">
        <Heading size="3xl" fontWeight="bold">Privacy policy</Heading>
        <Box>
          <Text>
            Copod is committed to protecting the privacy and security of your personal information. This
            Privacy Policy outlines how we collect, use, disclose, and protect the information we gather from users
            of
            Copod.
          </Text>
        </Box>
        <Heading size="2xl" fontWeight="bold">1. Information we collect</Heading>
        <Heading size="lg" fontWeight="semibold">a. Land</Heading>
        <Box>
          <Text>
            When registering with Copod, land owners are required to provide land details such as government-issued
            land title document for verification and blockchain registration purposes.
          </Text>
        </Box>
        <Heading size="lg" fontWeight="semibold">b. User</Heading>
        <Box>
          <Text>
            Users signing in to Copod are required to provide personal information such as email for account
            creation and authentication purposes, mobile location data for local land markets recommendations, and government-issued
            ID for verification and authentication purposes.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">2. How We User Your Information</Heading>
        <Heading size="lg" fontWeight="semibold">a. Land</Heading>
        <Box>
          <Text>
            We use land details provided during registration to identity and verify the property.
          </Text>
        </Box>
        <Heading size="lg" fontWeight="semibold">b. User</Heading>
        <Box>
          <Text>
            User information collected during sign-in is used for account creation, authentication, and communication
            purposes. We may also use this information to provide users with updates, recommendations services, and relevant service-related
            notifications.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">3. Information Sharing and Disclosure</Heading>
        <Heading size="lg" fontWeight="semibold">a. Internal Use</Heading>
        <Box>
          <Text>
            Copod may share information internally among our teams to provide and improve our services,
            facilitate communication, and ensure efficient operations.
          </Text>
        </Box>
        <Heading size="lg" fontWeight="semibold">b. Legal Compliance</Heading>
        <Box>
          <Text>
            We may disclose your information if required to do so by law or in response to valid legal requests, such as
            court orders or subpoenas. You will also get notification for any such activity.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">4. Data Security</Heading>
        <Box>
          <Text>
            Copod implements appropriate technical and organizational measures to safeguard the
            information we collect and maintain.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">5. Data Retention</Heading>
        <Box>
          <Text>
            We will retain your information only for as long as necessary to fulfill the purposes outlined in this
            Privacy Policy, unless a longer retention period is required or permitted by law.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">6. Your Rights</Heading>
        <Box>
          <Text>
            Users have the right to access, update, and correct their personal information held by Copod.
            Users also have the right to request deletion of their personal information <Link variant="underline" colorPalette="green.400" href="https://forms.gle/qcta7TCE5FYjyeb68">here</Link>, subject to legal requirements.
          </Text>
          <Text>
            Upon deletion, your account and all associated data will be permanently removed from our systems.
          </Text>
        </Box>
        <Heading size="3xl" fontWeight="bold">7. Changes to this Privacy Policy</Heading>
        <Box>
          <Text>
            Copod reserves the right to update or modify this Privacy Policy at any time. Any changes
            will
            be effective immediately upon posting the revised Privacy Policy on our website.
          </Text>
        </Box>
      </Container>
    </Box>
  )
}

export default Page
