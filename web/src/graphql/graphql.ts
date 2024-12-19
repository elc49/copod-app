/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
  UUID: { input: any; output: any; }
};

export type CreateOnboardingInput = {
  displayPictureUrl: Scalars['String']['input'];
  email: Scalars['String']['input'];
  supportdocUrl: Scalars['String']['input'];
  titleUrl: Scalars['String']['input'];
};

export type CreateUserInput = {
  email: Scalars['String']['input'];
  firstname: Scalars['String']['input'];
  lastname: Scalars['String']['input'];
  supportDocId: Scalars['UUID']['input'];
  supportDocVerification: Verification;
};

export type DisplayPicture = {
  __typename?: 'DisplayPicture';
  created_at: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  url: Scalars['String']['output'];
  verified: Verification;
};

export type GetUserLandsInput = {
  email: Scalars['String']['input'];
  verification: Verification;
};

export type Mutation = {
  __typename?: 'Mutation';
  chargeMpesa?: Maybe<Scalars['String']['output']>;
  createOnboarding: Onboarding;
  createUser: User;
  updateDisplayPictureVerificationById: DisplayPicture;
  updateTitleVerificationById: Title;
};


export type MutationChargeMpesaArgs = {
  input: PayWithMpesaInput;
};


export type MutationCreateOnboardingArgs = {
  input: CreateOnboardingInput;
};


export type MutationCreateUserArgs = {
  input: CreateUserInput;
};


export type MutationUpdateDisplayPictureVerificationByIdArgs = {
  input: UpdateDisplayPictureVerificationByIdInput;
};


export type MutationUpdateTitleVerificationByIdArgs = {
  input: UpdateTitleVerificationByIdInput;
};

export type Onboarding = {
  __typename?: 'Onboarding';
  created_at: Scalars['Time']['output'];
  displayPicture: DisplayPicture;
  displayPictureId: Scalars['UUID']['output'];
  email: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  supportDocId: Scalars['UUID']['output'];
  supportingDoc: SupportingDoc;
  title: Title;
  titleId: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  verification: Verification;
};

export enum PaidFor {
  NotPaid = 'NOT_PAID',
  Paid = 'PAID'
}

export type PayWithMpesaInput = {
  currency: Scalars['String']['input'];
  email: Scalars['String']['input'];
  paymentFor: Scalars['UUID']['input'];
  phone: Scalars['String']['input'];
  reason: PaymentReason;
};

export type Payment = {
  __typename?: 'Payment';
  created_at: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  onboarding?: Maybe<Onboarding>;
  onboarding_id: Scalars['UUID']['output'];
  reference_id: Scalars['String']['output'];
  status: Scalars['String']['output'];
  updated_at: Scalars['Time']['output'];
};

export enum PaymentReason {
  LandRegistry = 'LAND_REGISTRY'
}

export enum PaymentStatus {
  Failed = 'failed',
  Pending = 'pending',
  Success = 'success',
  Timeout = 'timeout'
}

export type PaymentUpdate = {
  __typename?: 'PaymentUpdate';
  email: Scalars['String']['output'];
  referenceId: Scalars['String']['output'];
  status: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  getDisplayPictureById: DisplayPicture;
  getOnboardingByEmail: Onboarding;
  getPaymentsByStatus: Array<Payment>;
  getSupportingDocById: SupportingDoc;
  getTitleById: Title;
  getUserLands: Array<Title>;
};


export type QueryGetDisplayPictureByIdArgs = {
  id: Scalars['UUID']['input'];
};


export type QueryGetOnboardingByEmailArgs = {
  email: Scalars['String']['input'];
};


export type QueryGetPaymentsByStatusArgs = {
  status: PaymentStatus;
};


export type QueryGetSupportingDocByIdArgs = {
  id: Scalars['UUID']['input'];
};


export type QueryGetTitleByIdArgs = {
  id: Scalars['UUID']['input'];
};


export type QueryGetUserLandsArgs = {
  input: GetUserLandsInput;
};

export type Subscription = {
  __typename?: 'Subscription';
  paymentUpdate: PaymentUpdate;
};


export type SubscriptionPaymentUpdateArgs = {
  email: Scalars['String']['input'];
};

export type SupportingDoc = {
  __typename?: 'SupportingDoc';
  created_at: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  url: Scalars['String']['output'];
  verified: Verification;
};

export type Title = {
  __typename?: 'Title';
  created_at: Scalars['Time']['output'];
  id: Scalars['UUID']['output'];
  support_doc_id: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  url: Scalars['String']['output'];
  verified: Verification;
};

export type UpdateDisplayPictureVerificationByIdInput = {
  displayPictureId: Scalars['UUID']['input'];
  verification: Verification;
};

export type UpdateTitleVerificationByIdInput = {
  titleId: Scalars['UUID']['input'];
  verification: Verification;
};

export type User = {
  __typename?: 'User';
  created_at: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  firstname?: Maybe<Scalars['String']['output']>;
  id: Scalars['UUID']['output'];
  lastname?: Maybe<Scalars['String']['output']>;
  updated_at: Scalars['Time']['output'];
};

export enum Verification {
  Onboarding = 'ONBOARDING',
  Rejected = 'REJECTED',
  Verified = 'VERIFIED'
}

export class TypedDocumentString<TResult, TVariables>
  extends String
  implements DocumentTypeDecoration<TResult, TVariables>
{
  __apiType?: DocumentTypeDecoration<TResult, TVariables>['__apiType'];

  constructor(private value: string, public __meta__?: Record<string, any> | undefined) {
    super(value);
  }

  toString(): string & DocumentTypeDecoration<TResult, TVariables> {
    return this.value;
  }
}
