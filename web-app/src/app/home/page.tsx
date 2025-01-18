import NextLink from "next/link"
import {
  Card,
  CardContent,
  CardHeader,
  CardFooter,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";

const reasons = ["Own land", "Search land", "Use land"]

export default function Page() {
  return (
    <div className="grid sm:grid-cols-1 md:grid-cols-2 h-dvh">
      <div className="bg-green-900 place-content-center h-dvh hidden md:block lg:block">
         <ul className="list-disc place-self-center space-y-4 md:px-4 text-white">
           {reasons.map((item: string, index: number) => (
             <div key={index} className="text-5xl md:text-6xl lg:text-7xl font-bold px-8">
               <li>
                 {item}
               </li>
             </div>
           ))}
         </ul>
      </div>
      <div className="h-full place-content-center">
        <div className="place-self-center">
          <Card className="w-[350px]">
            <CardHeader>
              <CardTitle>Log In</CardTitle>
            </CardHeader>
            <CardContent>
              <Button>Login with Google</Button>
            </CardContent>
            <CardFooter className="text-xs">
              <p>By signing in you agree with our <NextLink href="#" className="underline"> Privacy Policy</NextLink></p>
            </CardFooter>
          </Card>
        </div>
      </div>
    </div>
  )
}
