import * as yup from 'yup'
import groupBy from 'lodash/groupby'

const signInSchema = yup.object().shape({
  email: yup
    .string()
    .required('Email is required')
    .email('Email is not valid'),
  password: yup
    .string()
    .required('Password is required')
    .min(6, 'Password should be at least 6 characters long')
    .max(30, 'Password should be less then 30 characters long')
    .matches(/[A-Z]{1}/, 'Password should have at least 1 upper case letter')
    .matches(/[0-9]{1}/, 'Password should have at least 1 number')
    .matches(
      /[#?!@$%^&*-]{1}/,
      'Password should have at least 1 special character'
    )
})

export const validateSignInInput = async (input: {
  email?: string
  password?: string
}) => {
  const errors = await signInSchema
    .validate(input, { abortEarly: false })
    .then(() => undefined)
    .catch(err => err.inner)

  if (!errors) return {}

  const groupedErrors = groupBy(errors, 'path')
  for (let path in groupedErrors) {
    groupedErrors[path] = groupedErrors[path][0].errors[0]
  }
  return groupedErrors
}
