"use client";

import { AbsoluteCenter, Spinner } from "@chakra-ui/react";

function Loader() {
  return (
    <AbsoluteCenter>
      <Spinner size="md" />
    </AbsoluteCenter>
  );
}

export default Loader
