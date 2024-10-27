package com.lomolo.copodv2.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.layout.wrapContentWidth
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Close
import androidx.compose.material3.AlertDialogDefaults
import androidx.compose.material3.BasicAlertDialog
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import com.lomolo.copodv2.R
import com.lomolo.copodv2.ui.common.Avatar
import com.web3auth.core.types.UserInfo

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AccountDetails(
    modifier: Modifier = Modifier,
    setDialog: (Boolean) -> Unit,
    userInfo: UserInfo,
    signOut: () -> Unit,
) {
    BasicAlertDialog(onDismissRequest = {}) {
        Surface(
            modifier
                .wrapContentWidth()
                .wrapContentHeight(),
            shape = MaterialTheme.shapes.large,
            tonalElevation = AlertDialogDefaults.TonalElevation,
        ) {
            Column {
                IconButton(onClick = {
                    setDialog(false)
                }) {
                    Icon(
                        Icons.TwoTone.Close,
                        contentDescription = stringResource(R.string.close),
                    )
                }
                Column(
                    Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(16.dp)
                ) {
                    Column {
                        Row(
                            horizontalArrangement = Arrangement.spacedBy(8.dp),
                            verticalAlignment = Alignment.CenterVertically,
                        ) {
                           Avatar(
                               userInfo.profileImage,
                               userInfo.email
                           ) { setDialog(true) }
                            Column(
                                verticalArrangement = Arrangement.spacedBy(4.dp),
                            ) {
                                Text(
                                    userInfo.email,
                                    style = MaterialTheme.typography.bodyLarge,
                                    fontWeight = FontWeight.Bold,
                                    overflow = TextOverflow.Ellipsis,
                                )
                            }
                        }
                    }
                    TextButton(onClick = {
                        setDialog(false)
                        signOut()
                    }) {
                        Text(
                            stringResource(R.string.sign_out),
                            style = MaterialTheme.typography.bodyLarge,
                        )
                    }
                }
            }
        }
    }
}